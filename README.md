## Building device specific market app

Select a variant e.g. rpi or mp1, mp2

Select the corresponding __tag__, __architecture__ and __platform__, too

## Manual upgrade
    NAME=market
    VARIANT=rpi
    ctr -n system i pull -k ghcr.io/peramic/$NAME:${VARIANT}-latest
    systemctl stop $NAME
    ctr -n system c rm $NAME
    ctr -n system c create --label NAME=$NAME --label IS_ACTIVE=true --env LOGHOST=$NAME --with-ns network:/var/run/netns/$NAME ghcr.io/peramic/$NAME:${VARIANT}-latest $NAME
    systemctl start $NAME
