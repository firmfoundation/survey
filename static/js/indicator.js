function menuIndicator() {
    $(".content").empty();
    $(".content").append(`
        <div class="container-sm">
            <div id="left-col" class="cs-1 justify-content-center">
                <div class="select-survey form-floating mb-3 w-50"></div>

                <div class="form-floating mb-3 w-50 was-validated">
                    <input type="text" class="form-control form-control-sm" id="txt-indicator" value="" size="10" required>
                    <label for="txt-indicator" class="form-label">Indicator Name</label>
                    <div class="invalid-feedback">Indicator name required. </div>
                </div>

                <div class="form-floating mb-3 w-50 was-validated ">
                    <input type="text" class="form-control" id="indicator-weight" value="10" size="10" required>
                    <label for="indicator-weight" class="form-label">Indicator Weight</label>
                    <div class="invalid-feedback">Indicator weight need to be numeric. </div>
                </div>
            
                <div class="mt-3">
                    <button type="button" class="btn btn-warning btn-sm" id="btn-indicator-create" onClick="CreateIndicator()">Create Indicator</button>
                </div>            
            </div>
        </div>
    `)

    getSurveyList();
}

successGetSurvey = (json) => {
    listSurveys(json)
}

function CreateIndicator() {
    var select_value = $('.select-survey').find(":selected").val()
    if (select_value == "") {
        alert("select survey from the list!")
        return
    }
    var payload = {
        survey_id: select_value,
        name: $("#txt-indicator").val(),
        weight: Number($("#indicator-weight").val())
    };

    $.ajax({
        url: '/indicators',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(payload),
        success: function (response) {
           alert("Indicator created successfully!")
        },
        error: function (response) {
            alert("error")
        }
    })
}

function listSurveys(itm) {
    $(".select-survey").empty();
    var select = document.createElement('select');
    $(select).addClass('list-survey form-select')
    $(select).append(`<option value="" selected>select survey</option>`)
    $.each(itm, function( key, v ) {
        $(select).append(
            `<option value="${v.ID}">${v.name}</option>`
        )
    });
   

    $(".select-survey").append(select);
    $(".select-survey").append(
        `<label for="select-survey" class="form-label">Survey name</label>
        <div class="invalid-feedback">Survey name required. </div>`
    )
}

function getSurveyList() {
    $.ajax({
        url: '/surveys',
        type: 'GET',
        contentType: 'application/json',
        success: function (response) {
            var json = JSON.parse(response);
            successGetSurvey(json)
        },
        error: function (response) {
            
        }
    })
}



