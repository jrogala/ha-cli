Feature: Server configuration
  As an authenticated user I can retrieve the Home Assistant server info.

  Background:
    Given a running Home Assistant instance
    And an authenticated user

  Scenario: Get server config
    When I request the server config
    Then I should get a location name
    And I should get a version
    And I should get a timezone
