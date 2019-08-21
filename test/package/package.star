
def name():
    return "Test"


def commands(this, env):
    env.PYTHONPATH.append(this.root)
