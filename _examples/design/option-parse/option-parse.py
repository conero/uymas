


def option_parse(args, strict_option_list=None, map_data=None):
    '''
    选项解析算法实现(2019年12月20日 星期五)
    
    param: args  数组参数
    param: strict_option_list  严格数据列表
    param: map_data  option 映射键值

    见: _example/design/option-parse/option-parse.py
    '''
    # 选项队列
    options = []
    # 数据对象
    object_data = {}
    
    args_len = len(args)

    # 键存在标记
    last_key_mark = ''

    # 严格检查
    def strict_check_fn(copt):
        if copt and strict_option_list is not None and isinstance(strict_option_list, list):
            if not (copt in strict_option_list):
                raise Exception(f'The option <{copt}> is not allowed.')
    
    # 多选择检查
    def strict_check_list_fn(lists):
        if lists and  strict_option_list is not None and isinstance(strict_option_list, list):
            if isinstance(lists, list):
                for vlst in lists:
                    strict_check_fn(vlst)

    for i in range(0, args_len):
        arg = args[i].strip()
        arg_len = len(arg)

        if arg_len > 3:
            if arg[:2] == '--':
                arg = arg[2:]

                # --key=value
                if '=' in arg:
                    idx = arg.find('=')
                    cur_key = arg[:idx]
                    strict_check_fn(cur_key)

                    last_key_mark = ''

                    object_data[cur_key] = arg[idx+1:]
                    options.append(cur_key)
                    
                # --key
                else:
                    strict_check_fn(arg)
                    options.append(arg)
                    last_key_mark = arg
                
                continue

        if arg[:1] == '-':
            arg = arg[1:]
            tmp_list = list(arg)
            strict_check_list_fn(tmp_list)

            # list 合并
            options = options + tmp_list

            last_key_mark = options[len(options)-1]

        elif last_key_mark != '':
            # 已经存在值，认为时数组
            if last_key_mark in object_data:
                if isinstance(object_data[last_key_mark], list):
                    object_data[last_key_mark].append(arg)
                else:
                    tmp_value = object_data[last_key_mark]
                    object_data[last_key_mark] = []
                    object_data[last_key_mark].append(tmp_value)
                    object_data[last_key_mark].append(arg)

            else:
                object_data[last_key_mark] = arg

    return options, object_data





if __name__ == "__main__":
    # 数据测试
    opts, odd = option_parse(['-txf', '--table', 't1', 't2', 'tn', '--name="Joshua Conero"'])
    print(opts, odd)

    opts, odd = option_parse(['build', '--source', 't1.py', 'script.py', '--deepin'])
    print(opts, odd)

    opts, odd = option_parse(['build', '--source', 't1.py', 'script.py', '--deepin'], strict_option_list=['source'])
    print(opts, odd)

    # raise Exception('ssss')