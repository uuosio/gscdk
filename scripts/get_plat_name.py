import platform
#check the platform for linux, macos, windows
if platform.system() == "Linux":
    print("manylinux1_x86_64")
elif platform.system() == "Windows":
    print("win-amd64")
elif platform.system() == "Darwin":
    print("macosx_10_15_x86_64")
else:
    print("Unknown")
