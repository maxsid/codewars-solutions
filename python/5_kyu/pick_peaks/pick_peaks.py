def pick_peaks(arr):
    r, l, res = [0, 1, 2], len(arr), {"pos":[], "peaks":[]}
    while r[-1] < l:
        if arr[r[0]] < arr[r[1]] > arr[r[2]]:
            res['peaks'].append(arr[r[1]])
            res['pos'].append(r[1])
            r = [r[0] + x for x in range(2, 5)]
        elif arr[r[1]] == arr[r[2]]:
            r[2] += 1
        else:
            r = [r[0] + x for x in range(1, 4)]
    return res
