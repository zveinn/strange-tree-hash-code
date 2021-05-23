# random shit

# Notes
- Do not delete individual trees, the cleanup routine will be the only one removing links. This ensures minimum lock contention.

# Slice notes 
- Changing individual values does not trigger a copy of the slice.




# chain 
item - item - item - item
     - item - item - item

# int8 char based chain
root[int8]Item[int8]Item[int8]
