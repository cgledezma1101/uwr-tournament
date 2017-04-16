/*
	Responds to the asynchronous request of removing a goal from a player in a game

	Expects the following instance variables to be defined

	@param [Player] player Player that lost the goal
	@param [Integer] player_goals Total number of goals the player now has
	@param [Game] game The game where the goal was removed
	@param [Integer] team_goals Total number of goals the player's team now has
*/

$('.modal').modal('hide');

// First, find the player that scored and mark the goal
var players_goals = $('#goals-<%= @player.id %>');
players_goals.text(<%= @player_goals %>);

<% if @player_goals == 0 %>
	players_goals.parent().addClass('player-no-scores');
<% end %>

// Now increment the teams mark
<% if @player.team_id == @game.blue_team_id %>
	var selector = '.js-blue-goals';
<% else %>
	var selector = '.js-white-goals';
<% end %>

var teams_score = $(selector);
teams_score.text(<%= @team_goals %>);

addToFeeds("<%= t('score.remove_feed_text', player: @player.name, team: @player.team.name) %>")
