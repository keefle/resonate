#!/usr/bin/env sh

function cleanup() {
    cd /tmp/
    killall resonate 2> /dev/null

    fusermount -u dir1-resonate 2> /dev/null
    fusermount -u dir2-resonate 2> /dev/null
    rm -rf dir1-resonate 2> /dev/null
    rm -rf dir2-resonate 2> /dev/null
    rm -rf dir1 dir2 res-data  2> /dev/null
    echo "# Done cleanup"
    cd -
}

function test_resonate() {
    mkdir /tmp/res-data/
    cd /tmp/res-data
    if ( npm install local > /dev/null 2>&1 ); then
        echo "installed test data"
    else
        echo "failed to install test data"
        exit 1
    fi

    mkdir /tmp/dir1
    mkdir /tmp/dir2

    cd -
    bash -c './build/resonate -dir /tmp/dir1 -port 1234 -peer localhost:4321 2> /dev/null' &
    bash -c './build/resonate -dir /tmp/dir2 -port 4321 -peer localhost:1234 2> /dev/null' &

    sleep 1

    cp -a /tmp/res-data /tmp/dir1-resonate

    echo "# Started Checking diffrences"
    diff -qr /tmp/dir1-resonate /tmp/dir2-resonate
    diff -qr /tmp/dir1-resonate /tmp/dir1
    diff -qr /tmp/dir2-resonate /tmp/dir2
    diff -qr /tmp/dir1 /tmp/dir2
    diff -qr /tmp/dir1/res-data /tmp/res-data
    diff -qr /tmp/dir2/res-data /tmp/res-data
    echo "# Done Checking diffrences"
}

run_test() {
    cleanup
    test_resonate
    cleanup
}

run_test
