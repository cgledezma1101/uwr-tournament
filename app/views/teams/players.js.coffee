# Find the div where the players should be inserted
player_list = $(".players#<%= @team.id %>")
player_list.attr('id', '')

player_list.html('<%= j render "teams/players_checkboxes",
                               game: @game,
                               players: @players %>')
