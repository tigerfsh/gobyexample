import os 

if __name__ == "__main__":
    print("create dir")
    with open("all.txt") as f:
        i = 0
        while True:
            line = f.readline()
            if not line:
                break
            i +=1
            line = line.strip()
            s = line.lower()
            if i < 12:
                continue
            _dir_name = "-".join(s.split())
            dir_name = f"{i}-{_dir_name}"
            dir_name = dir_name.replace("/", "|")
            print(dir_name)
            ok = os.path.exists(f"/Users/a1/gobyexample/{dir_name}")
            if not ok:
                os.mkdir(dir_name)
            exists = os.path.exists(f"/Users/a1/gobyexample/{dir_name}/main.go")
            if not exists:
                f_path = f"/Users/a1/gobyexample/{dir_name}/main.go"
                ff = open(f_path, "w")
                ff.close()

