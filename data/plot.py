import numpy as np
import re
import matplotlib.pyplot as plt

def read_process_data(filename, samples):
    cpu = []
    mem = []

    cpu_count = 0
    mem_count = 0

    fd = open(filename, 'r')
    while cpu_count < samples or mem_count < samples:
        line = fd.readline()
        if line == '' or line is None:
            break

        n = re.search(r'CPU|MEM', line)
        if n is None:
            continue

        line = fd.readline()
        assert len(line) > 0
        fields = re.split(r'\s+', line)
        if n[0] == 'CPU':
            if cpu_count >= samples:
                continue
            cpu_count += 1
            cpu.append(float(fields[8]))
        elif n[0] == 'MEM':
            if mem_count >= samples:
                continue
            mem_count += 1
            mem.append(float(fields[8]))

    return cpu, mem


if __name__ == "__main__":
    gnet_cpu, gnet_mem = read_process_data('gnet.cpu', 90)
    go124_cpu, go124_mem = read_process_data('go124.cpu', 90)

    l = min(len(gnet_cpu), len(go124_cpu))
    gnet_cpu = gnet_cpu[0:l]
    gnet_mem = gnet_mem[0:l]
    go124_cpu = go124_cpu[0:l]
    go124_mem = go124_mem[0:l]

    fig, axs = plt.subplots(1, 2, figsize=(10, 4))
    x = np.arange(0, l)

    # for simplicity, we combine the two data sets into one
    axs[0].plot(x, gnet_cpu, label='gnet cpu%', color='blue', marker='o')
    axs[0].plot(x, go124_cpu, label='go124 cpu%', color='orange', marker='x')
    axs[0].set_title('CPU Usage Comparison')
    axs[0].set_xlabel('Tick (s)')
    axs[0].set_ylabel('CPU Usage (%)')
    axs[0].legend()
    axs[0].grid(True)

    axs[1].plot(x, gnet_mem, label='gnet mem%', color='blue', marker='o')
    axs[1].plot(x, go124_mem, label='go124 mem%', color='orange', marker='x')
    axs[1].set_title('MEM Usage Comparison')
    axs[1].set_xlabel('Tick (s)')
    axs[1].set_ylabel('MEM Usage (%)')
    axs[1].legend()
    axs[1].grid(True)

    plt.tight_layout()
    plt.show()
