(function (global) {

  var Stone = {
    Black: 1,
    White: 2
  };
  var GameResult = {
    Win: 1,
    Lose: 2,
    Draw: 3
  };
  var GameResultMessage = {};
  GameResultMessage[GameResult.Win] = {
    class: 'win',
    message: 'You win'
  };
  GameResultMessage[GameResult.Lose] = {
    class: 'lose',
    message: 'You lose'
  };
  GameResultMessage[GameResult.Draw] = {
    class: 'draw',
    message: 'Draw game'
  };

  function App(uri, output) {
    this.websocket = new WebSocket(uri);
    this.websocket.onopen = this.onOpen.bind(this);
    this.websocket.onclose = this.onClose.bind(this);
    this.websocket.onmessage = this.onMessage.bind(this);
    this.websocket.onerror = this.onError.bind(this);
    this.output = output;

    this.cells = [];
    this.board = document.getElementById('board');
    this.board.addEventListener('click', this.handleEvent.bind(this), false);

    this.message = document.getElementById('game-message');

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
      this.initBoard(msg.body.game);
    } else if (msg.type === 'nextTurn') {
      this.renderBoard(msg.body.game);
    } else if (msg.type === 'finish') {
      this.renderBoard(msg.body.game);
      this.renderMessage(msg.body.result);
    }
  }

  App.prototype.onError = function(evt) {
    console.log(evt)
  }

  App.prototype.onStartClick = function(evt) {
    var msg = JSON.stringify({
      type: 'start',
      body: {
        stone: Stone.Black
      }
    });
    this.websocket.send(msg);
  }

  App.prototype.handleEvent = function(evt) {
    if (evt.target.nodeName !== 'LI') {
      return;
    }
    this.onCellClick(evt);
  }

  App.prototype.onCellClick = function(evt) {
    var cell = evt.target;
    var x = cell.dataset.x;
    var y = cell.dataset.y;

    var msg = JSON.stringify({
      type: 'selectCell',
      body: {
        x: parseInt(x, 10),
        y: parseInt(y, 10)
      }
    });
    this.websocket.send(msg);
  }

  App.prototype.initBoard = function(game) {
    var i = 0;
    var x = 0;
    var y = 0;

    for (y = 0; game.board.size.height - 1 >= y; y++) {
      for (x = 0; game.board.size.width - 1 >= x; x++) {
        var cell = game.board.cells[i];
        var c = document.createElement('li');
        c.dataset.x = x;
        c.dataset.y = y;
        this.board.appendChild(c);
        this.cells.push(c);
        i++;
      }
    }
  }

  App.prototype.renderBoard = function(game) {
    var i = 0;

    for (; game.board.cells.length - 1 >= i; i++) {
      var c = this.cells[i];
      var cell = game.board.cells[i];

      if (cell.stone === Stone.Black) {
        c.innerText = 'B';
      } else if (cell.stone === Stone.White) {
        c.innerText = 'W';
      }
    }
  }

  App.prototype.renderMessage = function(status) {
    var message = GameResultMessage[status];
    this.message.innerText = message.message;
    this.message.classList.add(message.class);
  }

  global.App = App;

}(window));
