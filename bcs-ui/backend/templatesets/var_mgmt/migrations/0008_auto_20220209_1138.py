# Generated by Django 3.2.2 on 2022-02-09 03:38

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('variable', '0007_update_sys_vars'),
    ]

    operations = [
        migrations.AlterField(
            model_name='clustervariable',
            name='creator',
            field=models.CharField(max_length=64, verbose_name='创建者'),
        ),
        migrations.AlterField(
            model_name='clustervariable',
            name='updator',
            field=models.CharField(max_length=64, verbose_name='更新者'),
        ),
        migrations.AlterField(
            model_name='namespacevariable',
            name='creator',
            field=models.CharField(max_length=64, verbose_name='创建者'),
        ),
        migrations.AlterField(
            model_name='namespacevariable',
            name='updator',
            field=models.CharField(max_length=64, verbose_name='更新者'),
        ),
        migrations.AlterField(
            model_name='variable',
            name='creator',
            field=models.CharField(max_length=64, verbose_name='创建者'),
        ),
        migrations.AlterField(
            model_name='variable',
            name='updator',
            field=models.CharField(max_length=64, verbose_name='更新者'),
        ),
    ]