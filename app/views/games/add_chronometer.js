var chronometerPanel = $('.js-chronometers-panel');
var chronometers = chronometerPanel.children('.js-chronometers');

chronometers.append("<%= j render 'games/chronometer' %>");

$('.js-chronometer-start').on('click', chronometerStart);
$('.js-chronometer-stop').on('click', chronometerStop);
