class Club < ActiveRecord::Base
  has_many :user_clubs, dependent: :destroy
  has_many :users, through: :user_clubs

  has_many :club_admins, dependent: :destroy
  has_many :admins, through: :club_admins

  has_many :teams, dependent: :destroy

  validate :has_one_admin

  private
    def has_one_admin
      errors.add(t('club.errors.no_admin')) if self.admins.blank?
    end
end
