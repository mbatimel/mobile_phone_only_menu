UPDATE menu.dishes  SET favorite = true WHERE id = ANY($1::int[]);
-- Или альтернативный вариант:
-- UPDATE menu_dishes SET choice = true WHERE id = ANY($1::bigint[]);
-- UPDATE menu_dishes SET choice = false WHERE id != ALL($1::bigint[]);