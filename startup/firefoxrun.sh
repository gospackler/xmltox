function install_firefox {
  wget https://ftp.mozilla.org/pub/firefox/releases/48.0b5/linux-x86_64/en-US/firefox-48.0b5.tar.bz2 
  bzip2 -d firefox-*
  tar xf firefox-*
}

function start_with_xvfb {
  nohup xvfb-run -a ./firefox/firefox --profile $1 &	
}

function create_profile {
  mkdir -p $1
  echo 'user_pref("app.update.auto", false);' > $1/user.js
  echo 'user_pref("marionette.force-local", true);' >> $1/user.js
  echo 'user_pref("marionette.defaultPrefs.enabled", true);' >> $1/user.js
  echo 'user_pref("marionette.defaultPrefs.port",' $2');' >> $1/user.js
  start_with_xvfb $1
}

install_firefox
create_profile ~/.mozilla/profile/screenshot1 2828
create_profile ~/.mozilla/profile/screenshot2 2829
