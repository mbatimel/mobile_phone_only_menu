UPDATE menu.dishes 
SET name = $1 , category = $2 
WHERE id = $3;