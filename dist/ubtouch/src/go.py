import subprocess

def load():
    p1=subprocess.Popen(["/opt/click."+"ubuntu.com/jwstudy.anon/current/jwlib.bin"],stdout=subprocess.PIPE)
    print(p1.communicate()[0])
    return p1.communicate()[0]
