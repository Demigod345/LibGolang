#!/bin/bash

original_hash="US:$2y$10$bf0SYgc9JPOxSB1MG4eKVOL8/EhwcYcneLDzcVUkR/BT7Pmq4OwNu"

# Extract the part after the colon
hash_part=${original_hash:4}
echo "hash_part: " $hash_part

# Replace $2y$ with $2a$
new_hash="US:\$2a$"+"${hash_part:4}"

echo "Original hash: $original_hash"
echo "New hash: $new_hash"