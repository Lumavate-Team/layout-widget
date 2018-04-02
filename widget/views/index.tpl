<div class="starting">
	<div class="primary">
		{{range .primary}}

			<div class="name">
				{{.ComponentData.FirstName}} {{.ComponentData.LastName}}
			</div>

			<div class="info">
				{{.ComponentData.JobTitle}}<br>
				<a href="tel:{{.ComponentData.PhoneNumber}}">{{.ComponentData.PhoneNumber}}</a><br>
				<a href="mailto:{{.ComponentData.Email}}">{{.ComponentData.Email}}</a>
			</div>

			{{end}}
	</div>

	<HR COLOR="black"><br>

	<div class="secondary">
		{{range .secondary}}

			<div class="name">
				{{.ComponentData.FirstName}} {{.ComponentData.LastName}}
			</div>

			<div class="info">
				{{.ComponentData.JobTitle}}<br>
				<a href="tel:{{.ComponentData.PhoneNumber}}">{{.ComponentData.PhoneNumber}}</a><br>
				<a href="mailto:{{.ComponentData.Email}}">{{.ComponentData.Email}}</a>
			</div>

			{{end}}
	</div>
</div>