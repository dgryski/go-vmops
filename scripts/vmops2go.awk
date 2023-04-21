#!/usr/bin/awk

BEGIN {
    if (package == "") {
        package = "fsm";
    }
    if (var == "") {
        var = "Ops";
    }
    print "package", package;
    print;
    print "import \"github.com/dgryski/go-vmops\"\n";
    print "var", var, "= vmops.VM{";
}

# figure out the prefix we need to replace
/#ifndef.*LIBFSM_VMOPS_H/ {
    sub(/#ifndef /, "");
    sub(/LIBFSM.*/, "");
    prefix=$0;
    next;
}

# replace the comments
/^\t*\/\*/ {
    sub(/^\t*\/\*  */, "");
    sub(/ *\*\/$/, "");
    print "\t//", $0;
}

# replace the data lines
/^\t*{/ {
    # trim excess fluff from the line
    sub(/^\t*{/, "");
    sub(/},$/, "");
    # replace ',' and ' ' with hex to make splitting easier
    sub(/','/, "0x2c");
    sub(/' '/, "0x20");
    # extract the fields, splitting on comma
    split($0, arr, /, */);
    # fsm_opEOF => vmops.OpEOF
    sub(prefix "o", "vmops.O", arr[1]);
    # fsm_actionRET => vmops.ActionRET
    sub(prefix "a", "vmops.A", arr[3]);
    # undo the hex replacements
    if (arr[2] == "0x2c") { arr[2] = "','"; }
    if (arr[2] == "0x20") { arr[2] = "' '"; }
    # put back the line as a keyed initializer
    printf "\t{Op: %s, C: %s, Action: %s, Arg: %s},\n", arr[1], arr[2], arr[3], arr[4];
}

END {
    print "}";
}
