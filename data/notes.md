                                                                                           
```sql                                                                                        
SELECT                                                                                        
    *                                                                                         
FROM                                                                                          
    questions                                                                                 
ORDER BY created_time                                                                         
LIMIT 10                                                                                      
OFFSET 20;                                                                                    
```                                                                                           
    - the way of doing this will helps us to ignore offsets                                   
    - not using offset makes the query fast                                                   
                                                                                              
explain analyze select * from questions order by created_at desc offset 10000 limit 10;       
explain analyze select * from events order by id desc limit 10;                               
                                                                                              
```sql                                                                                        
postgres=# explain analyze select * from question order by created_at desc offset 10000000 lim
 Limit  (cost=19.57..19.57 rows=1 width=401) (actual time=0.028..0.029 rows=0 loops=1)        
   ->  Sort  (cost=19.09..19.57 rows=190 width=401) (actual time=0.025..0.026 rows=1 loops=1) 
         Sort Key: created_at DESC                                                            
         Sort Method: quicksort  Memory: 25kB                                                 
         ->  Seq Scan on events  (cost=0.00..11.90 rows=190 width=401) (actual time=0.014..0.0
 Planning Time: 0.165 ms                                                                      
 Execution Time: 0.075 ms                                                                     
```                                                                                           
see if we use offset we can see it will fetch the columns first and then limit it to 10. so th
we are doing this is by using the id ( unique)  as a offset to the value we used before and pa
using this

```sql
postgres=# explain analyze select * from events order by id desc limit 10;
 Limit  (cost=0.14..2.82 rows=10 width=401) (actual time=0.022..0.051 rows=1 loops=1)
   ->  Index Scan Backward using events_pkey on events  (cost=0.14..51.00 rows=190 width=401) 
 Planning Time: 0.145 ms
 Execution Time: 0.087 ms
