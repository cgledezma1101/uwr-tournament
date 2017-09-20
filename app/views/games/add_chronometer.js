var chronometerPanel = $('.js-chronometers-panel');
var chronometers = chronometerPanel.children('.js-chronometers');
var newChronometer = $("<%= j render 'games/chronometer' %>");

newChronometer.find('.js-chronometer-start').on('click', chronometerStart);
newChronometer.find('.js-chronometer-stop').on('click', chronometerStop);
newChronometer.find('.js-chronometer-remove').on('click', chronometerRemove);
newChronometer.find('.js-chronometer-external').on('click', chronometerExternal);

chronometers.append(newChronometer);
