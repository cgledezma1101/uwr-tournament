var playingChronometers = {};
var chronometerInterval = undefined;

var updateChronometers = function()
{
  $.each(playingChronometers, function(gameId, entries)
  {
    $.each(entries, function(index, entry)
    {
      var minutesInput = entry.chronometer.find('.js-chronometer-minutes');
      var secondsInput = entry.chronometer.find('.js-chronometer-seconds');

      var currentMinutes = +(minutesInput.val());
      var currentSeconds = +(secondsInput.val());

      var newSeconds = currentSeconds - 1;
      if (newSeconds === -1)
      {
        if (currentMinutes === 0)
        {
          entry.chronometer.find('.js-chronometer-stop').click();
        }
        else
        {
            minutesInput.val(currentMinutes - 1);
            secondsInput.val(59);
        }
      }
      else
      {
        secondsInput.val(newSeconds);
      }
    });
  });
}

var chronometerStart = function(event)
{
  event.preventDefault();
  var chronometer = $(this).parents('.js-chronometer');
  var name = chronometer.find('.js-chronometer-name');
  var minutes = chronometer.find('.js-chronometer-minutes');
  var seconds = chronometer.find('.js-chronometer-seconds');

  minutes.prop("readonly", "readonly");
  seconds.prop("readonly", "readonly");
  name.prop("readonly", "readonly");

  var gameId = chronometer.parents('.js-chronometers').data('game');
  var intervalEntry = {
    id: name.val(),
    chronometer: chronometer
  };
  if (!playingChronometers[gameId] || playingChronometers[gameId].length === 0)
  {
    playingChronometers[gameId] = [intervalEntry];
    chronometerInterval = window.setInterval(updateChronometers, 1000);
  }
  else
  {
    var isInArray = false;
    $.each(playingChronometers[gameId], function(index, entry)
    {
      if (entry.id == intervalEntry.id)
      {
        isInArray = true;
        return false;
      }

      return true;
    })

    if (!isInArray)
    {
      $.merge(playingChronometers[gameId], [intervalEntry]);
    }
  }
};

var chronometerStop = function(event)
{
  event.preventDefault();
  var chronometer = $(this).parents('.js-chronometer');
  var minutes = chronometer.find('.js-chronometer-minutes');
  var seconds = chronometer.find('.js-chronometer-seconds');
  var name = chronometer.find('.js-chronometer-name');

  minutes.prop("readonly", "");
  seconds.prop("readonly", "");
  name.prop("readonly", "");

  var gameId = chronometer.parents('.js-chronometers').data('game');
  var chronometerName = name.val();
  if (playingChronometers[gameId])
  {
    playingChronometers[gameId] = $.grep(playingChronometers[gameId], function(entry, index)
    {
      return entry.id !== chronometerName;
    });

    if (playingChronometers[gameId].length === 0)
    {
      window.clearInterval(chronometerInterval);
    }
  }
};

var chronometerRemove = function(event)
{
  event.preventDefault();
  var chronometer = $(this).parents('.js-chronometer');
  chronometer.find('.js-chronometer-stop').click();
  chronometer.remove();
}

$(document).ready(function()
{
  $('.js-chronometer-stopall').on('click', function(event)
  {
    event.preventDefault();
    var panel = $(this).parents('.js-chronometers-panel');
    panel.find('.js-chronometer-stop').click();
  });

  $('.js-chronometer-resumeall').on('click', function(event)
  {
    event.preventDefault();
    var panel = $(this).parents('.js-chronometers-panel');
    panel.find('.js-chronometer-start').click();
  });
});
