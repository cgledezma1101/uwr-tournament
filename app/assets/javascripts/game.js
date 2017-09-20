var registeredChronometers = {};
var chronometerInterval = undefined;
var activeChronometers = 0;
var scoreboardHandle = null;

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
			minutesInput.addClass('timer-low');
			secondsInput.addClass('timer-low');
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

	if (!chronometerId || (registeredChronometers[chronometerId] && !chronometer.hasClass('js-chronometer-playing')))
	{
		name.addClass('chronometer-error');
		return false;
	}

	name.removeClass('chronometer-error');

	minutes.prop("readonly", "readonly");
	seconds.prop("readonly", "readonly");
	name.prop("readonly", "readonly");

	if (activeChronometers === 0)
	{
		chronometerInterval = window.setInterval(updateChronometers, 1000);
	}

	if (!registeredChronometers[chronometerId])
	{
		chronometer.addClass('js-chronometer-playing');
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

	chronometer.removeClass('js-chronometer-playing');

	minutes.removeClass('timer-low');
	minutes.prop("readonly", "");

	seconds.removeClass('timer-low');
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

var chronometerExternal = function()
{
	$('.js-is-external').removeClass('js-is-external');
	$(this).addClass('.js-is-external');

	if (!scoreboardHandle)
	{
		scoreboardHandle = createScoreboardHandle();
	}
}

var createScoreboardHandle = function()
{
	return window.open('/games/external_scoreboard', '_blank', {
		location: 'no',
		menubar: 'no',
		status: 'no',
		titlebar: 'no'
	});
}

var addToFeeds = function(text)
{
	var event_form = $('#new_game_event').first();
	var input_field = event_form.find('.js-event-text');
	var old_text = input_field.val();
	input_field.val(text);
	event_form.submit();
	input_field.val(old_text);
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
