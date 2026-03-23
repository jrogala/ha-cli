Feature: Entity listing
  As an authenticated user I can list and filter entities.

  Background:
    Given a running Home Assistant instance
    And an authenticated user

  Scenario: List all entities
    Given entities exist in the system
    When I list all entities
    Then I should receive a list of entities
    And each entity should have an ID and state

  Scenario: Filter entities by domain
    Given entities exist in the system
    When I list entities with domain "light"
    Then all entities should belong to domain "light"

  Scenario: Search entities by name
    Given entities exist in the system
    When I list entities with search "kitchen"
    Then all entities should match search "kitchen"
