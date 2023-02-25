function menuQuestion() {
    $(".content").empty();
    $(".content").append(`
        <div class="container-sm">
            <div id="left-col" class="cs-1 was-validated justify-content-center">
                <div class="select-survey form-floating mb-3 w-50"></div>

                <div class="form-floating mb-3 w-50  was-validated">
                    <input type="text" class="form-control form-control-sm" id="question" value="" size="10" required>
                    <label for="question" class="form-label">Question</label>
                    <div class="invalid-feedback">Question field is required. </div>
                </div>

                <div class="select-indicator form-floating mb-3 w-50"></div>

                <div class="mt-3">
                    <button type="button" class="btn btn-warning btn-sm" id="btn-question-create" onClick="CreateQuestion()">Create Indicator</button>
                </div>            
            </div>
        </div>
    `)
    getSurveyList();
    getIndicatorList();
}

function CreateQuestion() {
    var select_survey_value = $('.select-survey').find(":selected").val()
    if (select_survey_value == "") {
        alert("select survey from the list!")
        return
    }
    var select_indicator_value = $('.select-indicator').find(":selected").val()
    if (select_indicator_value == "") {
        alert("select indicator from the list!")
        return
    }
    var payload = {
        survey_id: select_survey_value,
        indicator_id: select_indicator_value,
        question: $("#question").val(),
    };

    $.ajax({
        url: '/questions',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(payload),
        success: function (response) {
           alert("Question created successfully!")
        },
        error: function (response) {
            alert("error")
        }
    })
}

successGetIndicator = (json) => {
    listIndicators(json)
}

function listIndicators(itm) {
    $(".select-indicator").empty();
    var select = document.createElement('select');
    $(select).addClass('list-indicator form-select')
    $(select).append(`<option value="" selected>-Select Indicators-</option>`)
    $.each(itm, function( key, v ) {
        $(select).append(
            `<option value="${v.ID}">${v.name}</option>`
        )
    });
   

    $(".select-indicator").append(select);
    $(".select-indicator").append(
        `<label for="list-indicator" class="form-label">Indicator name</label>
        <div class="invalid-feedback">Indicator name required.</div>`
    )
}

function getIndicatorList() {
    $.ajax({
        url: '/indicators',
        type: 'GET',
        contentType: 'application/json',
        success: function (response) {
            var json = JSON.parse(response);
            successGetIndicator(json)
        },
        error: function (response) {
            
        }
    })
}





