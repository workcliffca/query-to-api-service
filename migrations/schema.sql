-- Create the API definitions table for Azure SQL Database
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[_api_definitions]') AND type in (N'U'))
BEGIN
    CREATE TABLE [dbo].[_api_definitions] (
        id INT IDENTITY(1,1) PRIMARY KEY,
        path NVARCHAR(255) NOT NULL UNIQUE,
        query NVARCHAR(MAX) NOT NULL,
        created_at DATETIME2 DEFAULT GETDATE(),
        updated_at DATETIME2 DEFAULT GETDATE(),
        is_active BIT DEFAULT 1
    );
END
GO

-- Create indexes for performance
IF NOT EXISTS (SELECT * FROM sys.indexes WHERE name = 'idx_api_definitions_path' AND object_id = OBJECT_ID('_api_definitions'))
BEGIN
    CREATE INDEX idx_api_definitions_path ON _api_definitions(path);
END
GO

IF NOT EXISTS (SELECT * FROM sys.indexes WHERE name = 'idx_api_definitions_active' AND object_id = OBJECT_ID('_api_definitions'))
BEGIN
    CREATE INDEX idx_api_definitions_active ON _api_definitions(is_active);
END
GO