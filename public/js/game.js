(function (global) {

  function App(uri, output) {
    this.websocket = new WebSocket(uri);
    this.websocket.onopen = this.onOpen.bind(this);
    this.websocket.onclose = this.onClose.bind(this);
    this.websocket.onmessage = this.onMessage.bind(this);
    this.websocket.onerror = this.onError.bind(this);
    this.output = output;

    this.startButton = document.getElementById('start');
    this.startButton.addEventListener('click', this.onStartClick.bind(this), false);
  }

  App.createApp = function(uri, output) {
    return new App(uri, output);
  }

  App.prototype.onOpen = function(evt) {
    console.log(evt)
  }

  App.prototype.onClose = function(evt) {
    console.log(evt)
  }

  App.prototype.onMessage = function(evt) {
    console.log(evt)
    var msg = JSON.parse(evt.data)
  }

  App.prototype.onError = function(evt) {
    console.log(evt)
  }

  App.prototype.onStartClick = function(evt) {
    var msg = JSON.stringify({
      type: 'start',
      body: null
    });
    this.websocket.send(msg);
  }

  global.App = App;

}(window));
