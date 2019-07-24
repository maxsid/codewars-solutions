def brain_luck(code, input):
    input, memory, data_pointer, code_pointer, output = list(input), dict(), 0, 0, ''
    brain_luck.code_pointer = 0
    case = lambda i: i == code[brain_luck.code_pointer]
    get_value = lambda: memory.get(data_pointer, 0)
    set_value = lambda v: memory.update({data_pointer: v % 256})

    def jump(step, start_symbol, last_symbol):
        skip = -1
        while skip != 0 or code[brain_luck.code_pointer] != last_symbol:
            if code[brain_luck.code_pointer] == start_symbol: skip += 1
            elif code[brain_luck.code_pointer] == last_symbol: skip -= 1
            brain_luck.code_pointer += step

    while brain_luck.code_pointer < len(code):
        if case('>'): data_pointer += 1
        elif case('<'): data_pointer -= 1
        elif case('+'): set_value(get_value() + 1)
        elif case('-'): set_value(get_value() - 1)
        elif case('.'): output += str(chr(get_value()))
        elif case(','): set_value(ord(input.pop(0)))
        elif case('[') and get_value() == 0: jump(1, '[', ']')
        elif case(']') and get_value() != 0: jump(-1, ']', '['); continue
        brain_luck.code_pointer += 1
    return output