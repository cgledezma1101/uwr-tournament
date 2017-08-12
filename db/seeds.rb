User.destroy_all
Club.destroy_all
Tournament.destroy_all

foo_users = (0..11).map{ |i| User.create(email: "foo#{i}@fooclub.com", name: "Foo#{i}", password: "foobarquux", password_confirmation: "foobarquux") }
bar_users = (0..11).map{ |i| User.create(email: "bar#{i}@barclub.com", name: "Bar#{i}", password: "foobarquux", password_confirmation: "foobarquux") }

foo_club = Club.create(name: 'FooClub')
bar_club = Club.create(name: 'BarClub')

ClubAdmin.create(club: foo_club, user: foo_users.first)
ClubAdmin.create(club: bar_club, user: bar_users.first)

foo_users.drop(1).each do |foo_user|
	UserClub.create(club: foo_club, user: foo_user)
end

bar_users.drop(1).each do |bar_user|
	UserClub.create(club: bar_club, user: bar_user)
end

foo_team = Team.create(name: 'FooTeam', club: foo_club)
foo_users.each_with_index do |foo_user, index|
	Player.create(user: foo_user, team: foo_team, number: index + 1, is_active: true)
end

bar_team = Team.create(name: 'BarTeam', club: bar_club)
bar_users.each_with_index do |bar_user, index|
	Player.create(user: bar_user, team: bar_team, number: index + 1, is_active: true)
end

tournament = Tournament.create(name: 'TheTournament', start_date: Date.today, end_date: Date.today + 3)
TournamentAdmin.create(tournament: tournament, user: foo_users.first)
TournamentAdmin.create(tournament: tournament, user: bar_users.first)

Stage.create(name: 'Stage0', tournament: tournament)

TournamentInvitation.create(tournament: tournament, club: foo_club)
TournamentInvitation.create(tournament: tournament, club: bar_club)

TournamentTeam.create(tournament: tournament, team: foo_team, password: 'footeam', password_confirmation: 'footeam')
TournamentTeam.create(tournament: tournament, team: bar_team, password: 'barteam', password_confirmation: 'barteam')