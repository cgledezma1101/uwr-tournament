class Club < ActiveRecord::Base
  has_many :user_clubs, dependent: :destroy
  has_many :users, through: :user_clubs

  has_many :club_admins, inverse_of: :club, dependent: :destroy
  has_many :admins, through: :club_admins, source: :user

  has_many :teams, dependent: :destroy

  validate :has_one_admin
  validates :name, presence: true, uniqueness: true

  private
    def has_one_admin
      errors.add(:admins, I18n.t('club.errors.no_admin')) if self.admins.blank?
    end
end
