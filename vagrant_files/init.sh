echo "
rdr pass inet proto tcp from any to 127.0.0.1 port 80 -> 127.0.0.1 port 8080
rdr pass inet proto tcp from any to 127.0.0.1 port 443 -> 127.0.0.1 port 4443
" | sudo pfctl -ef - >/dev/null 2>&1; echo "Add Port Forwarding (80 => 8080)\nAdd Port Forwarding (443 => 8443)"

