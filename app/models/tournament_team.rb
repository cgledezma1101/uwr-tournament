class TournamentTeam < ActiveRecord::Base
  belongs_to :tournament
  belongs_to :team

  validate :club_invited

  private

  # Makes sure that the club the team belongs to has been invited to this tournament
  def club_invited
    self.errors.add(:team, I18n.t('tournament.team_uninvited')) unless self.tournament.clubs.where{ id == my{ team.club.id } }.any?
  end
end
