import os
for f in os.listdir('./build'):
    if f.endswith('.tar.gz'):
        print(f)
        break
