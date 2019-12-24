import cmd_bin


def test_option_parse():
    # 数据测试
    opts, cmd, odd = cmd_bin.option_parse(['-txf', '--table', 't1', 't2', 'tn', '--name="Joshua Conero"'])
    print(opts, cmd, odd)

    opts, cmd, odd = cmd_bin.option_parse(['build', '--source', 't1.py', 'script.py', '--deepin'])
    print(opts, cmd, odd)

    # 严格检查开启，异常抛出检查
    # opts, cmd, odd = cmd_bin.option_parse(['build', '--source', 't1.py', 'script.py', '--deepin'], strict_option_list=['source'])
    # print(opts, cmd, odd)

    # raise Exception('ssss')


def test_cmd_cls():
    cmd = cmd_bin.Command()

    def cmd_version():
        print('v1.0.0')

    cmd.cmd_fn_register('version', cmd_version)
    cmd.opt_fn_register('v', cmd_version)

    cmd.args('-v')
    cmd.args('--any')

    cmd.run()


if __name__ == "__main__":
    # test_option_parse()
    test_cmd_cls()
