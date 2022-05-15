describe('The Register page', () => {





    it('test Twitter', () => {
        cy.visit('https://app2.abtasty.com/login?return_to=%2Fexperiments')

        cy.get('.FormButtonRow__buttonRow___3k8xJ > .Button__button___1rotk').should('be.disabled')

    })

    it('test Twitter', function () {
        cy.visit('https://www.instagram.com/?hl=fr')
        cy.get('.bIiDR').click()
        cy.get('.b_nGN').invoke('text').should('match', /^[0-9]*$/)
        cy.wait(3000)
        cy.get(':nth-child(1) > ._9GP1n > .f0n8F > ._2hvTZ').clear()
        cy.wait(3000)

        //cy.get('#password').type('truc')
        cy.get('.kEKum > :nth-child(3)').should('not.be.disabled')


    })


})