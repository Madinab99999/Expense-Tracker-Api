SQL-запросы для фильтрации и сортировки расходов:
1. Фильтр по прошлой неделе:
```sql
SELECT * 
FROM expense
WHERE user_id = 1
  AND date_expense >= NOW()::DATE - INTERVAL '7 days'
  AND date_expense < NOW()::DATE
```
2. Фильтр по прошлому месяцу:
```sql
SELECT * 
FROM expense
WHERE user_id = 1
  AND date_expense >= date_trunc('month', NOW()) - INTERVAL '1 month'
  AND date_expense < date_trunc('month', NOW());
```
3. Фильтр за последние 3 месяца:
```sql
SELECT * 
FROM expense
WHERE user_id = 1
  AND date_expense >= date_trunc('month', NOW()) - INTERVAL '3 months'
  AND date_expense < NOW();
```
4. Фильтр за этот год:
```sql
SELECT * 
FROM expense
WHERE user_id = 1
  AND date_expense >= date_trunc('year', NOW())
  AND date_expense < NOW();
```
5. Фильтр по пользовательскому диапазону дат:
```sql
SELECT * 
FROM expense
WHERE user_id = 1
  AND date_expense >= '2024-12-20'
  AND date_expense <= '2024-12-30';
```
6. Фильтр по минимальной и максимальной сумме:
```sql
SELECT * 
FROM expense
WHERE user_id = 1
  AND amount BETWEEN 50000 AND 30000;
```

SQL-запросы для получения статистики по расходам:
1. Пример запроса получения общей статистики по расходам:
```sql
SELECT 
    COALESCE(SUM(amount), 0) AS total_amount,
    COUNT(*) AS total_count,
    COALESCE(MAX(amount), 0) AS highest_amount,
    COALESCE(MIN(amount), 0) AS lowest_amount,
    COALESCE(AVG(amount), 0) AS average_amount
FROM expense
WHERE user_id = 1
  AND date_expense >= '2024-12-05'
  AND date_expense <= '2024-12-30';
```
2. Пример запроса получения статистики по категориям:
```sql
SELECT 
    category,
    COALESCE(SUM(amount), 0) AS total_amount,
    COUNT(*) AS count,
    COALESCE(MAX(amount), 0) AS highest_amount,
    COALESCE(MIN(amount), 0) AS lowest_amount,
    COALESCE(AVG(amount), 0) AS average_amount
FROM expense
WHERE user_id = 1
  AND date_expense >= '2024-12-05'
  AND date_expense <= '2024-12-30'
GROUP BY category
ORDER BY total_amount DESC;
```