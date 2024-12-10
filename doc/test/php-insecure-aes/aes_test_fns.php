<?php

function aesCbcTest()
{
    // print_r(openssl_get_cipher_methods());
    $data = 'Demo|十年生死两茫茫，不思量，自难忘|2018-07-01|㎡';
    $key = 'Y0^n4Dh47:dd7wOXyWJ9,jN-tv8jxY8i';
    $iv = '[K$jCYBej-vLVDQY';
    // $msg = openssl_encrypt($data, 'aes-256-cbc', $key, OPENSSL_ZERO_PADDING, $iv);
    $msg = openssl_encrypt($data, 'aes-256-cbc', $key, 0, $iv);
    if ($msg === false) {
        echo '加密失败, ' . openssl_error_string();
    } else {
        echo "密文：\n";
        echo $msg;
        echo "\n";

        // 还原
        $msg = 'E/6l1oVnotjo0gbovDloCuIa/CXdMysaOqQIWfZiYEpaJa+FjgFbESeVnZWoCeC7gNXsP/3Vd2OTB7X1b9jxE/PzThvC/fTiEuXVRm/J3o8=';
        $msg = openssl_decrypt($msg, 'aes-256-cbc', $key, 0, $iv);
        echo "还原明文：\n";
        echo $msg;
        echo "\n";
    }
}



// 执行
aesCbcTest();