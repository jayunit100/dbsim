### Bootstrap golang ### 

export PATH=$PATH:$HOME/bin/:$HOME/go/bin

echo "Modified path to include $HOME/go/bin."

function install_prerequisites {
        if [[ ! -f ~/bin/gimme ]] ; then 
                mkdir -p  ~/bin/
                mv ~/bin/gimme /tmp/
                curl -sL -o ~/bin/gimme https://raw.githubusercontent.com/travis-ci/gimme/master/gimme
                chmod +x ~/bin/gimme
        else
                echo "gimme found"
        fi
}
function install_go {
        echo "Installing Golang 1.11"
        eval "$(GIMME_GO_VERSION=1.11 gimme)"
        go version
}

install_prerequisites
install_go
