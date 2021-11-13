#!/bin/bash
if [[ "X$ABSTRUSE_BRANCH" == "X" || "X$ABSTRUSE_BRANCH" == "Xmaster" ]];
then
    echo -n -e ""
else
    echo -n -e "-$ABSTRUSE_BRANCH"
fi;