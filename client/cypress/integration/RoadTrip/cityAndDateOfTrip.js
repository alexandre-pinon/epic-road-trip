describe('The cities and dates of the trip', () => {

    it('the start city is the same on the startEnd Page', () => {
        cy.visit('http://localhost:3000/')
        cy.fixture('tripInformations').then((tripInformations) => {
            //cy.log(tripInformations.startCity)
            cy.get('[data-testid="searchBar"]')
                .type(tripInformations.startCity+'{enter}')
                .type('{enter}')
            cy.wait(250)

            cy.location('href').should('include', '/startEndTrip')

            cy.get('[data-testid="StartCity"]').should('have.value', tripInformations.startCity)
        })
    })









})