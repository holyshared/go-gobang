(function (global) {

  function App(uri, output) {
    this.websocket = new WebSocket(uri);
    this.websocket.onopen = this.onOpen.bind(this);
    this.websocket.onclose = this.onClose.bind(this);
    this.websocket.onmessage = this.onMessage.bind(this);
    this.websocket.onerror = this.onError.bind(this);
    this.output = output;

    this.board = document.getElementById('board');

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
  
    if (msg.type === 'start') {
      this.board.setAttribute('class', 'display');
      this.startButton.setAttribute('class', 'hidden');
      renderBoard(this.board, msg.body.game);
    }
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

  function renderBoard(el, game) {
    var i = 0;
    var x = 0;
    var y = 0;

    for (y = 0; game.board.size.height - 1 >= y; y++) {
      for (x = 0; game.board.size.width - 1 >= x; x++) {
        var cell = game.board.cells[i];
        var c = document.createElement('li');

        if (cell.stone === 1) {
          c.innerText = 'B';
        } else if (cell.stone === 2) {
          c.innerText = 'W';
        }
        el.appendChild(c);
        i++;
      }
    }
  }

  global.App = App;

}(window));
