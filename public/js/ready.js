(function () {

  window.addEventListener("DOMContentLoaded", function() {
    var output = document.getElementById('game');
    var app = App.createApp('ws://localhost:3000/ws', output);

  }, false);

}());
