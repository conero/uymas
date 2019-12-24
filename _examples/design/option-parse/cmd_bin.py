# coding=utf-8
# version=python > 3.6
import sys


def option_parse(args, strict_option_list=None, map_data=None):
    """
    选项解析算法实现(2019年12月20日 星期五)/函数式
    :param args: 数组参数
    :param strict_option_list: 严格数据列表
    :param map_data: option 映射键值
    :return: (options, cmds, object_data)
    """
    # 选项队列
    options = []
    # 数据对象
    object_data = {}
    # 命令
    cmds = []

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
        if lists and strict_option_list is not None and isinstance(strict_option_list, list):
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

                    object_data[cur_key] = arg[idx + 1:]
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

            last_key_mark = options[len(options) - 1]

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

        # 将命令行放到命令列表内
        elif len(options) == 0 and len(object_data) == 0:
            cmds.append(arg)

    return options, cmds, object_data


class App:
    def __init__(self, options=None, odata=None):
        if odata is None:
            odata = {}
        if options is None:
            options = []
        self._options = options
        self._odata = odata

    def check_opts(self, *options):
        """
        选择存在检测
        :param options:
        :return:
        """
        check = False
        for opt in options:
            if opt in self._options:
                check = True
                break
        return check

    def getOpts(self, *options, def_value=None):
        """
        获取值
        :param options:
        :param def_value:
        :return:
        """
        for opt in options:
            if opt in self._odata:
                def_value = self._odata[opt]
                break
        return def_value

    def data(self):
        return self._odata


class Command:
    """
    命令行程序实现
    """

    def __init__(self, **params):
        self._options = []  # 选项
        self._odata = {}  # 选项数据
        self._args = []  # 参数
        self._map_data = {}
        self._cmds = []
        self.app = None
        self._cmd_fun_callbacks = {}  # 命令回调函数
        self._opt_fun_callback = {}
        self._router_mark = False  # 路由成功标识
        self._callfn_empty_index = None

        if len(sys.argv) > 0:
            self._args = sys.argv[1:]

        if params:
            if 'map_data' in params:
                self._map_data = params['map_data']

    def empty_fn_register(self, callfn):
        if callable(callfn):
            self._callfn_empty_index = callfn
        else:
            raise Exception(f'empty_fn_register: the callfn is not valid callback')
        return self

    def cmd_fn_register(self, cmd, callfn):
        if callable(callfn):
            self._cmd_fun_callbacks[cmd] = callfn
        else:
            raise Exception(f'cmd_fn_register: the callfn is not valid callback. cmd: {cmd}')
        return self

    def opt_fn_register(self, cmd, callfn):
        if callable(callfn):
            if isinstance(cmd, list):
                for c in cmd:
                    self._opt_fun_callback[c] = callfn
            else:
                self._opt_fun_callback[cmd] = callfn
        else:
            raise Exception(f'opt_fn_register: the callfn is not valid callback. cmd: {cmd}')
        return self

    def args(self, *v_args):
        """
        修改参数
        :param v_args:
        :return:
        """
        self._args = []

        for arg in v_args:
            self._args.append(arg)

    def _do_router(self):
        """
        项目路由
        :return:
        """
        self._router_mark = False
        not_empty = len(self._options)
        not_empty = True if not_empty == 0 and not_empty == len(self._odata) else False

        for cmd in self._cmd_fun_callbacks:
            if cmd in self._cmds:
                self._cmd_fun_callbacks[cmd]()
                self._router_mark = True
                break

        if not self._router_mark:
            for opt in self._opt_fun_callback:
                if opt in self._options:
                    self._opt_fun_callback[opt]()
                    self._router_mark = True
                    break

        if not self._router_mark:
            if not_empty and callable(self._callfn_empty_index):
                self._callfn_empty_index()

            elif not_empty:
                print(f'Cannot handler the commands: {self._cmds} or options: {self._options}')

    def run(self):
        """
        运行命令程序
        :return:
        """
        self._options, self._cmds, self._odata = option_parse(self._args, map_data=self._map_data)
        self.app = App(self._options, self._odata)
        self._do_router()
