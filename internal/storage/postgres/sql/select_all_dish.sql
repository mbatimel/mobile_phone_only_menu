SELECT id, name, category, favorite 
FROM menu.dishes 
WHERE create_at::date = $1::date
ORDER BY id; 