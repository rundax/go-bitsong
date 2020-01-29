---
home: true
actionText: Get Started  →
actionLink: /guide/
footer: MIT Licensed | Copyright @ 2020 bitsong
---

<div class="features">
  <div class="feature">
    <h2>Tokenize your productions</h2>
    <p>Tokenize your productions, receive new investments for your new productions and share your rewards with all your fans.</p>
  </div>
  <div class="feature">
    <h2>Cosmos-SDK based</h2>
    <p>BitSong is a new music platform, which is built using <a href="https://cosmos.network" target="_blank">Cosmos-SDK</a></b> and the distributed <a href="https://ipfs.io/" target="_blank">IPFS</a></b> filesystem.</p>
  </div>
  <div class="feature">
    <h2>Blockchain Distribution</h2>
    <p>Your music will be distributed via BitSong on all its clients. For each track listened, you will receive a real-time <b>BTSG reward</b>.</p>
  </div>
  <div class="feature">
    <h2>Immutable streams</h2>
    <p>Thanks to blockchain technology, each stream will be recorded and impossible to modify.</p>
  </div>
  <div class="feature">
    <h2>Payments in real time</h2>
    <p>When one of your tunes is listened to, <b><i>you will be rewarded in real time in BTSG, and your fans will receive a portion of your rewards.</i></b></p>
  </div>
  <div class="feature">
    <h2>Distributied Governance</h2>
    <p>Distributed governance means that there is no element that can make decisions independently.</p>
  </div>
</div>

## Try BitSong Network
Currently tested for macOS and Linux

### 1. Install go 1.13+
``` bash
wget https://dl.google.com/go/go1.13.6.linux-amd64.tar.gz
sudo tar -xvzf go1.13.6.linux-amd64.tar.gz
sudo mv go /usr/local
cat <<EOF >> ~/.profile  
export GOPATH=$HOME/go  
export GO111MODULE=on  
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin  
EOF

source ~/.profile
```

### 2. Clone go-bitsong
``` bash
git clone https://github.com/bitsongofficial/go-bitsong.git && cd go-bitsong
```

### 3. Checkout the latest version
``` bash
git checkout v0.3.0
```

### 4. Compile go-bitsong
``` bash
make install

# go-bitsong version
bitsongd version --long
```