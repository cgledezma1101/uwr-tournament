class Tournament < ActiveRecord::Base
  has_many :tournament_admins, dependent: :destroy
  has_many :admins, through: :tournament_admins, source: :user

  has_many :tournament_invitations, dependent: :destroy
  has_many :clubs, through: :tournament_invitations

  has_many :tournament_teams, dependent: :destroy
  has_many :teams, through: :tournament_teams

  validates :name, presence: true
  validates :start_date, presence: true
  validates :end_date, presence: true
  validate :date_range

  private

  # Validates that the start date of the tournament is before the end date
  def date_range
    self.errors.add(:start_date, I18n.t('tournament.date_mismatch')) if start_date > end_date
  end
end
