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
    this.selectedStone = Stone.Black;
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

    this.startMenu = document.getElementById('startMenu');
    this.startMenu.addEventListener('click', this.handleEvent.bind(this), false);
    this.startMenus = document.querySelectorAll('#startMenu li a');

    this.startButton = document.getElementById('start');
    this.startButton.addEventListener('click', this.handleEvent.bind(this), false);

    this.retryMenu = document.getElementById('retryMenu');

    this.retryButton = document.getElementById('retry');
    this.retryButton.addEventListener('click', this.handleEvent.bind(this), false);

    this.gameStarted = false;
    this.initilalized = false;
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
      this.startMenu.setAttribute('class', 'hidden');
      this.render(msg.body.game);
      this.gameStarted = true;
    } else if (msg.type === 'nextTurn') {
      this.render(msg.body.game);
    } else if (msg.type === 'finish') {
      this.gameStarted = false;
      this.render(msg.body.game);
      this.renderMessage(msg.body.result);
      this.retryMenu.classList.remove('hidden');
    }
  }

  App.prototype.onError = function(evt) {
    console.log(evt)
  }

  App.prototype.handleEvent = function(evt) {
    var target = evt.target;

    evt.preventDefault();
    evt.stopPropagation();

    if (target.nodeName === 'LI' && target.dataset.type === 'cell' && this.gameStarted) {
      this.onCellClick(evt);
    } else if (target.nodeName === 'A' && target.dataset.stone) {
      this.onStoneClick(evt);
    } else if (target.nodeName === 'BUTTON' && target.id === 'start') {
      this.onStartClick(evt);
    } else if (target.nodeName === 'BUTTON' && target.id === 'retry') {
      this.onRetryClick(evt);
    }
  }

  App.prototype.onStartClick = function(evt) {
    var msg = JSON.stringify({
      type: 'start',
      body: {
        stone: this.selectedStone
      }
    });
    this.websocket.send(msg);
  }

  App.prototype.onRetryClick = function(evt) {
    var msg = JSON.stringify({
      type: 'start',
      body: {
        stone: this.selectedStone
      }
    });
    this.websocket.send(msg);
    this.clearMessage();
    this.retryMenu.classList.add('hidden');
  }

  App.prototype.onStoneClick = function(evt) {
    var target = evt.target;
    var stone = target.dataset.stone;

    this.startMenus.forEach(function (menu) {
      menu.classList.remove('selected');
    });

    target.classList.add('selected');
    this.selectedStone = parseInt(stone, 10);
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

  App.prototype.render = function(game) {
    if (this.initilalized) {
      return this.renderBoard(game);
    }
    this.initBoard(game);
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
        c.dataset.type = 'cell';
        this.board.appendChild(c);
        this.cells.push(c);
        i++;
      }
    }
    this.initilalized = true;
  }

  App.prototype.renderBoard = function(game) {
    var i = 0;

    for (; game.board.cells.length - 1 >= i; i++) {
      var c = this.cells[i];
      var cell = game.board.cells[i];

      if (cell.stone === Stone.Black) {
        c.innerText = '⚫️';
      } else if (cell.stone === Stone.White) {
        c.innerText = '⚪️';
      } else {
        c.innerText = '';
      }
    }
  }

  App.prototype.renderMessage = function(status) {
    var message = GameResultMessage[status];
    this.message.innerText = message.message;
    this.message.classList.add(message.class);
  }

  App.prototype.clearMessage = function() {
    [ 'win', 'lose', 'draw' ].forEach(function (className) {
      this.message.classList.remove(className);
    }, this);
    this.message.innerText = '';
  }

  global.App = App;

}(window));
