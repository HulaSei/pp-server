-- Historical security/audit rows predate centralized redaction. Remove
-- credentials and message bodies while retaining IP/User-Agent risk metadata.
UPDATE system_logs
SET content = CASE type
    WHEN 10 THEN (
        content::jsonb ||
        '{"to":"[REDACTED]","subject":"[REDACTED]","content":{"redacted":true},"template":"[REDACTED]"}'::jsonb
    )::text
    WHEN 11 THEN (
        content::jsonb ||
        '{"to":"[REDACTED]","subject":"[REDACTED]","content":{"redacted":true},"template":"[REDACTED]"}'::jsonb
    )::text
    WHEN 20 THEN (
        content::jsonb ||
        '{"token":"[REDACTED]"}'::jsonb
    )::text
    WHEN 31 THEN (
        content::jsonb ||
        '{"identifier":"[REDACTED]"}'::jsonb
    )::text
    ELSE content
END
WHERE type IN (10, 11, 20, 31);

-- Invalid historical retention values must not turn the next daily cleanup
-- into a broad future-dated delete.
UPDATE "system"
SET value = '7', type = 'int64'
WHERE category = 'log'
  AND key = 'ClearDays'
  AND CASE
      WHEN value ~ '^[0-9]+$' THEN value::numeric < 1 OR value::numeric > 3650
      ELSE TRUE
  END;
