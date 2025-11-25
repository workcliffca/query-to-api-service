-- Create the API definitions table for PostgreSQL
CREATE TABLE IF NOT EXISTS _api_definitions (
    id SERIAL PRIMARY KEY,
    path VARCHAR(255) NOT NULL UNIQUE,
    query TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT true
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_api_definitions_path ON _api_definitions(path);
CREATE INDEX IF NOT EXISTS idx_api_definitions_active ON _api_definitions(is_active);

-- Create a function to automatically update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create trigger to automatically update updated_at
DROP TRIGGER IF EXISTS update_api_definitions_updated_at ON _api_definitions;
CREATE TRIGGER update_api_definitions_updated_at
    BEFORE UPDATE ON _api_definitions
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();