var registeredChronometers = {};
var chronometerInterval = undefined;
var activeChronometers = 0;

var updateChronometers = function()
{
  $.each(registeredChronometers, function(chronometerId, chronometer)
  {
    if (!chronometer)
    {
      return true;
    }

    var minutesInput = chronometer.find('.js-chronometer-minutes');
    var secondsInput = chronometer.find('.js-chronometer-seconds');

    var currentMinutes = +(minutesInput.val());
    var currentSeconds = +(secondsInput.val());

    --currentSeconds;

    if (currentMinutes === 0 && currentSeconds <= 15)
    {
      minutesInput.css('background-color', '#FFFD60');
      secondsInput.css('background-color', '#FFFD60');
    }

    if (currentSeconds === -1)
    {
      if (currentMinutes === 0)
      {
        chronometer.find('.js-chronometer-stop').click();
      }
      else
      {
          minutesInput.val(--currentMinutes);
          secondsInput.val(59);
      }
    }
    else
    {
      secondsInput.val(currentSeconds);
    }

    return true;
  });
}

var chronometerStart = function()
{
  var chronometer = $(this).parents('.js-chronometer');
  var name = chronometer.find('.js-chronometer-name');
  var minutes = chronometer.find('.js-chronometer-minutes');
  var seconds = chronometer.find('.js-chronometer-seconds');

  chronometerId = name.val();

  if (!chronometerId || registeredChronometers[chronometerId])
  {
    name.css('border-color', 'rgb(233, 102, 102)');
    name.css('box-shadow', '0px 1px 1px rgba(0, 0, 0, 0.075) inset, 0px 0px 8px rgba(233, 102, 102, 0.6)');
    return false;
  }

  name.css('border-color', '');
  name.css('box-shadow', '');

  minutes.prop("readonly", "readonly");
  seconds.prop("readonly", "readonly");
  name.prop("readonly", "readonly");

  if (activeChronometers === 0)
  {
    chronometerInterval = window.setInterval(updateChronometers, 1000);
  }

  if (!registeredChronometers[chronometerId])
  {
    registeredChronometers[chronometerId] = chronometer;
    ++activeChronometers;
  }

  // Prevent default behaviour and stop event propagation
  return false;
};

var chronometerStop = function()
{
  var chronometer = $(this).parents('.js-chronometer');
  var minutes = chronometer.find('.js-chronometer-minutes');
  var seconds = chronometer.find('.js-chronometer-seconds');
  var name = chronometer.find('.js-chronometer-name');

  minutes.css('background-color', '');
  minutes.prop("readonly", "");

  seconds.css('background-color', '');
  seconds.prop("readonly", "");

  name.prop("readonly", "");

  var chronometerId = name.val();
  if (registeredChronometers[chronometerId])
  {
    registeredChronometers[chronometerId] = undefined
    --activeChronometers;

    if (activeChronometers === 0)
    {
      window.clearInterval(chronometerInterval);
    }
  }

  // Prevent default behaviour and stop event propagation
  return false;
};

var chronometerRemove = function()
{
  var chronometer = $(this).parents('.js-chronometer');
  chronometer.find('.js-chronometer-stop').click();
  chronometer.remove();

  // Prevent default behaviour and stop event propagation
  return false;
}

$(document).ready(function()
{
  $('.js-chronometer-stopall').on('click', function()
  {
    var panel = $(this).parents('.js-chronometers-panel');
    panel.find('.js-chronometer-stop').click();

    // Prevent default behaviour and stop event propagation
    return false;
  });

  $('.js-chronometer-resumeall').on('click', function()
  {
    var panel = $(this).parents('.js-chronometers-panel');
    panel.find('.js-chronometer-start').click();

    // Prevent default behaviour and stop event propagation
    return false;
  });
});
