-- Historical security/audit rows predate centralized redaction. Remove raw
-- bearer credentials and personal request metadata already persisted there.
UPDATE system_logs
SET content = CASE type
    WHEN 10 THEN JSON_SET(
        content,
        '$.to', '[REDACTED]',
        '$.subject', '[REDACTED]',
        '$.content', JSON_OBJECT('redacted', TRUE),
        '$.template', '[REDACTED]'
    )
    WHEN 11 THEN JSON_SET(
        content,
        '$.to', '[REDACTED]',
        '$.subject', '[REDACTED]',
        '$.content', JSON_OBJECT('redacted', TRUE),
        '$.template', '[REDACTED]'
    )
    WHEN 20 THEN JSON_SET(
        content,
        '$.token', '[REDACTED]',
        '$.user_agent', '[REDACTED]',
        '$.client_ip', '[REDACTED]'
    )
    WHEN 30 THEN JSON_SET(
        content,
        '$.login_ip', '[REDACTED]',
        '$.user_agent', '[REDACTED]'
    )
    WHEN 31 THEN JSON_SET(
        content,
        '$.identifier', '[REDACTED]',
        '$.register_ip', '[REDACTED]',
        '$.user_agent', '[REDACTED]'
    )
    ELSE content
END
WHERE type IN (10, 11, 20, 30, 31)
  AND JSON_VALID(content);

-- Invalid historical retention values must not turn the next daily cleanup
-- into a broad future-dated delete.
UPDATE `system`
SET value = '7', type = 'int64'
WHERE category = 'log'
  AND `key` = 'ClearDays'
  AND (CAST(value AS SIGNED) < 1 OR CAST(value AS SIGNED) > 3650);
