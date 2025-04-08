## alvk

**https://education.vk.company/program/kurs-algoritmy-structury-dannyh - Алгоритмы и структуры данных**


### Single Linked List

#### Reverse
```
prev -> nil    
 | 5 | -> | 4 | -> | 3 | -> | 2 | -> | 1 |
  /
current  

node -> current.next | 4, 3, 2, 1, nil
current.next -> prev | nil, 5...
prev -> current
current -> next

      prev
    /
 | 5 | -> | 4 | -> | 3 | -> | 2 | -> | 1  |
           /
       current 
```

