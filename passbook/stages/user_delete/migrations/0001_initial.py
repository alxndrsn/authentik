# Generated by Django 3.0.5 on 2020-05-15 19:59

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("passbook_flows", "0001_initial"),
    ]

    operations = [
        migrations.CreateModel(
            name="UserDeleteStage",
            fields=[
                (
                    "stage_ptr",
                    models.OneToOneField(
                        auto_created=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        parent_link=True,
                        primary_key=True,
                        serialize=False,
                        to="passbook_flows.Stage",
                    ),
                ),
            ],
            options={
                "verbose_name": "User Delete Stage",
                "verbose_name_plural": "User Delete Stages",
            },
            bases=("passbook_flows.stage",),
        ),
    ]
