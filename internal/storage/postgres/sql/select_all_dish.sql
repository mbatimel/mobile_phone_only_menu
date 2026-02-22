SELECT id, name, category, favorite
FROM menu.dishes
WHERE  create_at >= $1
  AND create_at < $1 + INTERVAL '1 day'
ORDER BY id;