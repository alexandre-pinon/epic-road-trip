describe('The Register page', () => {





    it('test Twitter', () => {
        cy.visit('https://app2.abtasty.com/login?return_to=%2Fexperiments')
        cy.get('#email').type('hehooo')
        cy.wait(3000)
        cy.get('#email').clear()

        cy.get('.Layout__leftColumn___2R6OU').click()
        //cy.get('[data-testid="password"] > .Input__label___1E1EQ').type('test')

        //cy.get('.Input__errorMessage___k6Dtz').should('have.text', 'Please enter a valid email')

        //cy.get('#email').clear().should('have.value', '');
        cy.get('.FormButtonRow__buttonRow___3k8xJ > .Button__button___1rotk').should('be.disabled')

    })







/*
    it('test Twitter', function () {
        cy.visit('https://www.instagram.com/?hl=fr')
        cy.get('.bIiDR').click()
        cy.get(':nth-child(1) > ._9GP1n > .f0n8F > ._2hvTZ').typ('hehoooo')
        cy.wait(3000)
        cy.get(':nth-child(1) > ._9GP1n > .f0n8F > ._2hvTZ').clear()
        cy.wait(3000)

        //cy.get('#password').type('truc')
        cy.get('.kEKum > :nth-child(3)').should('not.be.disabled')


    })

 */


})