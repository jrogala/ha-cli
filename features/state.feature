Feature: Entity state
  As an authenticated user I can get the state of a specific entity.

  Background:
    Given a running Home Assistant instance
    And an authenticated user

  Scenario: Get entity state
    Given entity "light.kitchen" exists
    When I get the state of "light.kitchen"
    Then I should get entity details
    And the entity ID should be "light.kitchen"
    And the entity should have a state value

  Scenario: Get state of non-existent entity
    When I get the state of "light.nonexistent"
    Then it should fail with "404"
